import { Injectable } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { ProcessInvoiceFraudDTO } from '../dto/process-invoice-fraud.dto';
import { Account, FraudReason, InvoiceStatus } from '../../../generated/prisma';
import { FraudAggregateSpecification } from './specifications/fraud-aggregate.specification';

@Injectable()
export class FraudService {
  constructor(
    private prismaService: PrismaService,
    private fraudAggregateSpec: FraudAggregateSpecification,
  ) {}

  async processInvoice(processInvoiceFraudDto: ProcessInvoiceFraudDTO) {
    const { invoice_id, account_id, amount } = processInvoiceFraudDto;

    const foundInvoice = await this.prismaService.invoice.findUnique({
      where: {
        id: invoice_id,
      },
    });

    if (foundInvoice) {
      throw new Error('Invoice has already been processed');
    }

    const account = await this.prismaService.account.upsert({
      where: {
        id: account_id,
      },
      update: {},
      create: {
        id: account_id,
      },
    });

    const fraudResult = await this.fraudAggregateSpec.detectFraud({
      account,
      amount,
      invoiceId: invoice_id,
    });

    const invoice = await this.prismaService.invoice.create({
      data: {
        id: invoice_id,
        accountId: account.id,
        amount,
        ...(fraudResult.hasFraud && {
          fraudHistory: {
            create: {
              reason: fraudResult.reason!,
              description: fraudResult.description,
            },
          },
        }),
        status: fraudResult.hasFraud
          ? InvoiceStatus.REJECTED
          : InvoiceStatus.APPROVED,
      },
    });

    return {
      invoice,
      fraudResult,
    };
  }
}
import type React from "react"
import { Inter } from "next/font/google"
import "./globals.css"
import { Header } from "@/components/header"
import { Toaster } from "react-hot-toast";

const inter = Inter({ subsets: ["latin"] })

export const metadata = {
  title: "Payment Gateway",
  description: "Gateway de pagamentos completo",
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="pt-BR">
      <body className={`${inter.className} bg-[#1a2332] min-h-screen`}>
        <Toaster position="top-right" />
        <Header />
        <main className="container mx-auto py-4">{children}</main>
        <footer className="py-4 text-center text-sm text-gray-400">
          © 2025 Payment Gateway.
        </footer>
      </body>
    </html>
  )
}
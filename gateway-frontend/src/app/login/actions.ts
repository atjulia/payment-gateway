"use server";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export async function loginAction(_: any, formData: FormData) {
  const apiKey = formData.get("apiKey") as string;

  const response = await fetch("http://localhost:8080/accounts", {
    headers: {
      "X-API-KEY": apiKey,
    },
  });
  if (!response.ok) {
    return { error: "API Key inválida. Verifique e tente novamente." };
  }

  const cookieStore = await cookies();
  cookieStore.set("apiKey", apiKey);

  redirect("/invoices");
}

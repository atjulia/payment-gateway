"use client";

import { InfoIcon } from "lucide-react";
import { Alert, AlertDescription, AlertTitle } from "../../components/ui/alert";
import { Button } from "../../components/ui/button";
import { Input } from "../../components/ui/input";
import { useEffect, useActionState } from "react";
import toast from "react-hot-toast";
import { loginAction } from "./actions";

export function AuthForm() {
  const initialState = { error: null as string | null };
  const [state, formAction] = useActionState(loginAction, initialState);

  useEffect(() => {
    if (state?.error) {
      toast.error(state.error);
    }
  }, [state]);

  return (
    <form className="space-y-4" action={formAction}>
      <div className="space-y-2">
        <label htmlFor="apiKey" className="text-sm text-gray-300">
          API Key
        </label>
        <div className="flex gap-2">
          <Input
            id="apiKey"
            placeholder="Digite sua API Key"
            className="bg-[#2a3749] border-gray-700 text-white placeholder-gray-400"
            name="apiKey"
          />
          <Button
            type="submit"
            className="bg-indigo-600 hover:bg-indigo-700 text-white"
          >
            →
          </Button>
        </div>
      </div>

      <Alert className="bg-[#2a3749] border-gray-700 mt-4">
        <InfoIcon className="h-4 w-4 text-blue-400" />
        <AlertTitle className="text-gray-200">Como obter uma API Key?</AlertTitle>
        <AlertDescription className="text-gray-400">
          Para obter sua API Key, você precisa criar uma conta.
          Caso não tenha um usuário, entre em contato com nosso suporte.
        </AlertDescription>
      </Alert>
    </form>
  );
}

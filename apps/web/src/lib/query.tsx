import { queryOptions, mutationOptions } from "@tanstack/react-query";
import { product, productInsertSchema } from "../types/index";
import type z from "zod";
const APIROUTE = "http://localhost:8080";

export const productsQueryOptions = queryOptions({
  queryKey: ["product"],
  queryFn: async (): Promise<z.infer<typeof product>[]> => {
    const response = await fetch(`${APIROUTE}/products`);
    if (!response.ok) {
      throw new Error("Network response was not ok");
    }
    return response.json();
  },
});

export const createProductOptions = mutationOptions({
  mutationKey: ["products"],
  mutationFn: async (data: z.infer<typeof productInsertSchema>) => {
    const res = await fetch(`${APIROUTE}/products`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });
    if (!res.ok) {
      throw new Error("Failed to create product");
    }

    return await res.json();
  },
  // onSuccess: (createdProduct) => {
  //       queryClient.setQueryData(
  //         ["products"],
  //         (old: z.infer<typeof product>[] | undefined) => {
  //           if (!old) return [createdProduct];
  //           return [...old, createdProduct];
  //         }
  //       );
  //     },
});

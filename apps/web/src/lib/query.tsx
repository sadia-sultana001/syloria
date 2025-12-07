import { queryOptions, mutationOptions, QueryClient } from "@tanstack/react-query";
import { product, productInsertSchema } from "../types/index";
import type z from "zod";
const APIROUTE = "http://localhost:8080";

export const productsQueryOptions = queryOptions({
  queryKey: ["products"],
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
 onMutate : async (variables , ctx)=>{
        // Cancel any outgoing refetches
      await ctx.client.cancelQueries({ queryKey: ['products'] })


       // Snapshot the previous producs
    const previousProducts = ctx.client.getQueryData(['products'])

    // Optimistically update to the new product
    ctx.client.setQueryData(['products'], (old: z.infer<typeof product>[] ) => [...old, variables])
      

    return { previousProducts }
 },


  // If the mutation fails,
  // use the result returned from onMutate to roll back
  onError: (err, newProduct, onMutateResult, context) => {
    context.client.setQueryData(['products'], onMutateResult?.previousProducts)
  },

   // Always refetch after error or success:
  onSettled: (data, error, variables, onMutateResult, context) =>
    context.client.invalidateQueries({ queryKey: ['products'] }),

})
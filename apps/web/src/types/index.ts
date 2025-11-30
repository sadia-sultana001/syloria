import * as z from "zod"; 

export const product = z.object({ 
  id: z.int(),
  title: z.string(),
  description: z.string(),
  price: z.float64(),
  imageUrl: z.string()
});

export const productInsertSchema = product.omit({id: true})
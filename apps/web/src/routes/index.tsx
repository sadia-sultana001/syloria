import { createFileRoute } from '@tanstack/react-router'
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Button } from '@/components/ui/button'

import ProductForm from '@/components/form/form.product'
import { productsQueryOptions } from '@/lib/query'
import { useSuspenseQuery } from '@tanstack/react-query'
import ProductCard from '@/components/product-card'
export const Route = createFileRoute('/')({
  component: App,
loader: ({ context: { queryClient } }) =>
    queryClient.ensureQueryData(productsQueryOptions),
})



function App() {
  const productsquery = useSuspenseQuery(productsQueryOptions)
  const products = productsquery.data
  return (
<main className='container mx-auto py-10'>
    <h3 className='text-center text-3xl font-semibold'>WELCOME TO SIMPLE STORE</h3>
   
   {/* product form in dialog */}
   <div className='flex justify-end w-full'>
   <Dialog>
      <form>
        <DialogTrigger asChild>
          <Button variant="outline">NEW PRODUCT</Button>
        </DialogTrigger>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>ADD PRODUCT</DialogTitle>
            <DialogDescription>
              CREATE A NEW PRODUCT......
            </DialogDescription>
          </DialogHeader>
          <ProductForm/>
          <DialogFooter>
            <DialogClose asChild className='w-full'>
              <Button variant="outline">Cancel</Button>
            </DialogClose>

          </DialogFooter>
        </DialogContent>
      </form>
    </Dialog>
</div>

 <h3 className='text-center text-2xl font-semibold mt-10 mb-5'> PRODUCT LIST</h3>

<div className='grid lg:grid-cols-4 md:grid-cols-2 grid-cols-1 gap-6  place-items-center'>
  {products.map((prdocut)=>(
    <ProductCard {...prdocut} key={prdocut.id} />
  ))}
</div>
</main>  
)
}

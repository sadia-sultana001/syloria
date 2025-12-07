import ProductCard from '@/components/product-card';
import { productByIdQueryOprions } from '@/lib/query'
import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'
;

export const Route = createFileRoute('/products/$productId')({
  component: RouteComponent,
  loader: async ({ params , context }) => {
    return  await context.queryClient.ensureQueryData(productByIdQueryOprions(params))
  },
})

function RouteComponent() {
  const {productId} = Route.useParams()
   const productquery = useSuspenseQuery(productByIdQueryOprions({productId}))
   const productData = productquery.data
   console.log("🚀 ~ RouteComponent ~ productData:", productData)
  return(
    <section className='flex h-screen justify-center items-center'>
           
           <ProductCard  {...productData} key={productData.id}/>
    </section>
  )
}

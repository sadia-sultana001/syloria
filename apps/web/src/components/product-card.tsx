
import React from 'react'
import { Card, CardAction, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from './ui/card'
import { Button } from './ui/button'
import type z from 'zod'
import type { product } from '@/types'

function ProductCard({id, description ,title , imageUrl ,price}:z.infer<typeof product>) {
  return (
     <Card className="w-full max-w-sm">
      <CardHeader>
        <CardTitle>{title}</CardTitle>
        <CardDescription>
         {description}
        </CardDescription>
        <CardAction>
          <Button variant="link">Details</Button>
        </CardAction>
      </CardHeader>
      <CardContent>
       <img className='w-full object-cover h-36 rounded-2xl' src={imageUrl} alt={title}/>
      </CardContent>
      <CardFooter className="flex-col gap-2">
        <p>Price: ${price}</p>
      </CardFooter>
    </Card>
  )
}

export default ProductCard



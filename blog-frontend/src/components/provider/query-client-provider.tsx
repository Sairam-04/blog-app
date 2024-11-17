import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import {ReactQueryDevtools} from "@tanstack/react-query-devtools"
import { useState } from "react"

type QueryClientProviderProps = {
    children: React.ReactNode
}

export function QueryClientWrapper({children}: QueryClientProviderProps){
    const [queryClient] = useState(()=> new QueryClient())

    return(
        <QueryClientProvider client={queryClient}>
            {children}
            <ReactQueryDevtools />
        </QueryClientProvider>
    )

}
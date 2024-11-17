import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import { routeTree } from './routeTree.gen.ts'
import { createRouter, RouterProvider } from '@tanstack/react-router'
import { QueryClientWrapper } from './components/provider/query-client-provider.tsx'

const router = createRouter({routeTree})
declare module '@tanstack/react-router'{
  interface Register{
    router: typeof router
  }
}


createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <QueryClientWrapper>
      <RouterProvider router={router}>
      </RouterProvider>
    </QueryClientWrapper>
  </StrictMode>,
)

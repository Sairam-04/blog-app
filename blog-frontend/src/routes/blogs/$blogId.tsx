import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/blogs/$blogId')({
  component: RouteComponent,
})

function RouteComponent() {
  const {blogId} = Route.useParams()
  return(
    <div>

      This is blog no : {blogId}
    </div>
  )
}

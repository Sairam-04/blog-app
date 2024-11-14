import Blog from '@/components/molecules/blog'
import { BLOGS } from '@/data/data'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/blogs/$blogId')({
  component: RouteComponent,
})

function RouteComponent() {
  const {blogId} = Route.useParams()
  const blog = BLOGS.filter((ele) => ele.id === blogId)
  const blogData = blog[0]
  return(
      <Blog {...blogData} />
  )
}

import BlogCard from '@/components/molecules/blog-card'
import { BLOGS } from '@/data/data'
import { createFileRoute, Link } from '@tanstack/react-router'

export const Route = createFileRoute('/blogs/')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className='grid grid-cols-3 gap-5 py-4'>
      {
        BLOGS.map((ele)=>{
          return(
            <Link to={`/blogs/${ele.id}`}>
              <BlogCard {...ele} key={ele.id} />
            </Link>
          )
        })
      }
    </div>
  )
}

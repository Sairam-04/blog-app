import { Card, CardContent, CardDescription, CardTitle } from "../ui/card";
import { blog } from "@/types/blog";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { formatDate } from "@/utils/utils";

export default function BlogCard({title, description, created_at, updated_at, thumbnail, content, name, id}: blog) {
    console.log(title, description, created_at, updated_at, thumbnail, content, name, id)
  return (
    <Card className="w-full max-w-sm rounded-lg overflow-hidden shadow-lg border-none">
        <img 
            src={thumbnail}
            alt="blogImage"
            className="w-full h-48 object-cover"
            width="400"
            height="200"
            style={{ aspectRatio: "400/200", objectFit: "cover" }}
        />
        <CardContent className="p-6 space-y-4">
        <div>
          <CardTitle className="text-xl font-semibold">{title}</CardTitle>
          <CardDescription className="text-gray-500">{description}</CardDescription>
        </div>
        <div className="flex justify-between items-center">
            <div className="flex gap-4 items-center">
                <Avatar>
                    <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
                    <AvatarFallback>CN</AvatarFallback>
                </Avatar>
                <div>{name}</div>
            </div>
          <div className="text-gray-500 text-sm">
            {formatDate(created_at)}
          </div>
        </div>
      </CardContent>
    </Card>
  )
}

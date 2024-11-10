import { blog } from "@/types/blog";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { formatDate } from "@/utils/utils";
import DOMPurify from 'dompurify';

export default function Blog({ title, description, created_at, updated_at, thumbnail, content, name, id }: blog) {
    console.log(id, updated_at, description)
    const sanitizedContent = DOMPurify.sanitize(content);
    return (
        <div className="w-full flex flex-col gap-5 pt-10">
            <div className="text-3xl font-bold">
                {title}
            </div>
            <div className="flex gap-6 items-center text-lg">
                <div className="flex gap-4 items-center">
                    <Avatar>
                        <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
                        <AvatarFallback>CN</AvatarFallback>
                    </Avatar>
                    <div className="text-base">{name}</div>
                </div>
                <div className="text-base">{formatDate(created_at)}</div>
            </div>
            <div className="w-full">
                <img
                    src={thumbnail}
                    alt="blogImage"
                    className="w-full h-96 object-cover"
                    style={{ objectFit: "cover" }}
                />
            </div>
            <div
                className="prose dark:prose-invert bg-transparent"
                dangerouslySetInnerHTML={{ __html: sanitizedContent }}
            />
        </div>
    )
}


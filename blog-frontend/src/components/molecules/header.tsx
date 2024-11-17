import { Link } from "@tanstack/react-router";
import { Button } from "../ui/button";
import ThemeButton from "./themeButton";
import { useGloablDrawerStore } from "./slider/store";

export default function Header() {
  const { setOpenGlobalDrawer, setGlobalDrawerType} = useGloablDrawerStore()
  return (
    <div className="sticky top-0 h-16 w-full py-4 px-2 backdrop-blur-3xl">
      <div className=" w-4/5 mx-auto flex justify-between items-center">
        <div className="logo">
          <div className="text-2xl font-semibold">Logo</div>
        </div>
        <NavItems />
        <div className="flex gap-3">
          <Button variant="secondary"
            onClick = {()=>{
              setOpenGlobalDrawer(true)
              setGlobalDrawerType("LOGIN_USER")
            }}
          >Login</Button>
          <ThemeButton />
        </div>
      </div>
    </div>
  );
}


const NavItems = () => {
  return(
    <div className="flex items-center gap-4">
      <Link to="/" className="cursor-pointer font-medium">Home</Link>
      <Link to="/" className="cursor-pointer font-medium">About</Link>
      <Link to="/blogs" className="cursor-pointer font-medium">Blog</Link>
      <Link to="/" className="cursor-pointer font-medium">PostBlog</Link>
    </div>
  )
}
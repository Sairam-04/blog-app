import { createRootRoute, Outlet } from "@tanstack/react-router";

export const Route = createRootRoute({
    component: ()=>{
        return(
            <div>
                Root Component
                <Outlet />
            </div>
        )
    }
})
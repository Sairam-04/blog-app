import { createRootRoute, Outlet } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/router-devtools";
import { ThemeProvider } from "../components/molecules/theme-provider";
import Header from "../components/molecules/header";

export const Route = createRootRoute({
    component: ()=>{
        return(
            <ThemeProvider>
                <div className="w-full">
                    <Header />
                    <div className="w-4/5 mx-auto">
                        <Outlet />
                    </div>
                    <TanStackRouterDevtools />
                </div>
            </ThemeProvider>
        )
    }
})
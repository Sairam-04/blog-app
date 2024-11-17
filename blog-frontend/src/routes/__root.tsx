import { createRootRoute, Outlet } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/router-devtools";
import { ThemeProvider } from "../components/provider/theme-provider";
import Header from "../components/molecules/header";
import { useGloablDrawerStore } from "@/components/molecules/slider/store";
import GlobalDrawer from "@/components/molecules/slider";

export const Route = createRootRoute({
    component: RootLayout
})

function RootLayout(){
    const {openGlobalDrawer} = useGloablDrawerStore()
    return(
        <ThemeProvider>
            <div className="w-full">
                <Header />
                <div className="w-4/5 mx-auto">
                    <Outlet />
                </div>
                {openGlobalDrawer && <GlobalDrawer />}
                <TanStackRouterDevtools />
            </div>
        </ThemeProvider>
    )
}
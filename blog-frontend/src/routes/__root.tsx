import { createRootRoute, Outlet } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/router-devtools";
import { ThemeProvider } from "../components/provider/theme-provider";
import Header from "../components/molecules/header";
import { useGloablDrawerStore } from "@/components/molecules/slider/store";
import GlobalDrawer from "@/components/molecules/slider";
import Gradient from "@/components/molecules/gradient";

export const Route = createRootRoute({
  component: RootLayout,
});

function RootLayout() {
  const { openGlobalDrawer } = useGloablDrawerStore();
  return (
    <div>
      <ThemeProvider>
        <div className="w-full relative">
          <Gradient />
          <div className="sticky w-4/5 mx-auto top-0 h-16 py-4 px-2 backdrop-blur-3xl">
            <Header />
          </div>
          <div className="w-4/5 mx-auto">
            <Outlet />
          </div>
          {openGlobalDrawer && <GlobalDrawer />}
          <TanStackRouterDevtools />
        </div>
      </ThemeProvider>
    </div>
  );
}

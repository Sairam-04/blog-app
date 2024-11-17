import { Drawer, DrawerContent } from "@/components/ui/drawer";
// import { Suspense } from "react";
import { useGloablDrawerStore } from "./store";
import { cn } from "@/lib/utils";
import LoginDrawer from "../drawers/login-drawer";
import RegisterDrawer from "../drawers/register.drawer";


export default function GlobalDrawer(){
    const {
        openGlobalDrawer,
        GlobalDrawerType,
        setOpenGlobalDrawer,
    } = useGloablDrawerStore()
    return(
        <Drawer
            direction="right"
            open={openGlobalDrawer}
            onClose={() =>{
                setOpenGlobalDrawer(false)
            }}
            onOpenChange={open =>{
                setOpenGlobalDrawer(open)
            }}
            handleOnly
        >
            <DrawerContent
                className={cn(
					"h-screen w-1/2 rounded-bl-md rounded-tl-md p-6 pb-4 pt-2 transition-width",
				)}
				// withCloseable={false}
				onClick={e => e.stopPropagation()}
            >
                {/* <Suspense fallback={<>Loading...</>}> */}   
                    {
                        GlobalDrawerType === "LOGIN_USER" && (<LoginDrawer  onClose={setOpenGlobalDrawer} />)
                    }
                    {
                        GlobalDrawerType === "REGISTER_USER" && (<RegisterDrawer />)
                    }
                    
                    {/* <RegistrationDrawerContent /> */}
                {/* </Suspense> */}
            </DrawerContent>
        </Drawer>
    )
}
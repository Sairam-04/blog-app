import { Button } from "@/components/ui/button";

export default function LoginDrawer({onClose} : {onClose: (open: boolean) => void}){
    return(
        <div className="w-full">
            <div>Login Drawer</div>
            <Button variant="default" onClick={()=>onClose(false)} >Cancel</Button>
        </div>
    )
}
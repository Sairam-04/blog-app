import { useContext } from "react"
import { ThemeProviderContext } from "../components/provider/theme-provider"

export const useTheme = () =>{
    const context = useContext(ThemeProviderContext)
    if (context === undefined){
        throw new Error("useTheme must be within a ThemeProvider")
    }
    return context
}
import { MONTHS } from "@/constants/constants"

export function formatDate(date: string):string {
    const newDate = new Date(date)
    const month = newDate.getMonth()
    const day = newDate.getDate()
    const year = newDate.getFullYear()
    const formattedString = MONTHS[month] + " " + day + ", "+ year
    return formattedString
}
import { create } from "zustand"

type initialStateType = {
    openGlobalDrawer: boolean
    setOpenGlobalDrawer: (state: boolean) => void

    GlobalDrawerType:
    | "REGISTER_USER"
    | "LOGIN_USER"
    | "CREATE_BLOG"

    setGlobalDrawerType: (
        state:
            | "REGISTER_USER"
            | "LOGIN_USER"
            | "CREATE_BLOG"
    ) => void
}

export const useGloablDrawerStore = create<initialStateType>(set => ({
    openGlobalDrawer: false,
    setOpenGlobalDrawer: (state: boolean) => set({ openGlobalDrawer: state }),
    GlobalDrawerType: "REGISTER_USER",
    setGlobalDrawerType: (
        state:
            | "REGISTER_USER"
            | "LOGIN_USER"
            | "CREATE_BLOG"
    ) => set({ GlobalDrawerType: state })
}))
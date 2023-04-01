import { atom } from 'jotai'

type Alert = {
  alert: "success" | "danger"
  message: string
}


export const alertAtom = atom<Alert>({ alert: "success", message: "" })

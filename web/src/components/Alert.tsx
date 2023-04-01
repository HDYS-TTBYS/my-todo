import { useAtom } from 'jotai'
import React, { FC } from 'react'
import { alertAtom } from '../store'



const Alert: FC = () => {
  const [alert] = useAtom(alertAtom);
  return (
    <div hidden={alert.message === ""} className={"alert alert-" + alert.alert} role="alert">
      {alert.message}
    </div>
  )
}

export default Alert

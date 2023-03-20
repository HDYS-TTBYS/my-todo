import React, { FC } from 'react'

interface Props {
  title: string
}

const Since: FC<Props> = ({ title }) => {
  return (
    <div className="d-flex align-items-center p-3 my-3 text-white bg-purple rounded shadow-sm">
      <div className="lh-1">
        <h1 className="h6 mb-0 text-white lh-1">My Todo</h1>
        <small>Since 2023</small>
      </div>
      <small className='px-3'>{title}</small>
    </div>
  )
}

export default Since

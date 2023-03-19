import React, { FC } from 'react'
import { useParams } from 'react-router-dom'
import NavBar from '../components/NavBar';
import { useQueryTodo } from '../hooks/useQueryTodo'



const Detail: FC = () => {
  const { id } = useParams();
  const { data, isLoading, error } = useQueryTodo(Number(id))

  if (isLoading) return <div> 'Loading...'</div>

  if (error) return <div>'error has occurred: ' + error</div>

  return (
    <>
      <NavBar />
      <div className='container'>
        <div className='d-flex'>
          タイトル: {data?.title}
        </div>
        <div className='d-flex'>
          担当者: @{data?.assagin_person}
        </div>
        <div className='d-flex'>
          状態: {data?.is_complete ? "完了" : "未完了"}
        </div>
        <div className='d-flex'>
          説明: {data?.description}
        </div>
        <div className='d-flex'>
          作成日: {new Date(data?.created_at! * 1000).toLocaleDateString() + "|" + new Date(data?.created_at! * 1000).toLocaleTimeString()}
        </div>
        <div className='d-flex'>
          更新日:  {new Date(data?.updated_at! * 1000).toLocaleDateString() + "|" + new Date(data?.updated_at! * 1000).toLocaleTimeString()}
        </div>
      </div>
    </>
  )
}

export default Detail

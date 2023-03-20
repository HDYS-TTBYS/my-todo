import React, { FC } from 'react'
import { Link, useNavigate } from 'react-router-dom';
import { useParams } from 'react-router-dom'
import Loading from '../components/Loading';
import NavBar from '../components/NavBar';
import Since from '../components/Since';
import { useQueryTodo } from '../hooks/useQueryTodo'



const Detail: FC = () => {
  const { id } = useParams<"id">();
  const router = useNavigate();
  const { data, isLoading, error } = useQueryTodo(Number(id))

  if (isLoading) return <Loading />
  if (error) router("/")

  return (
    <>
      <NavBar />
      <div className='container'>
        <Since title='詳細ページ' />
        <div className='flex mt-5'>
          <div className='d-flex mb-2'>
            タイトル: {data?.title}
          </div>
          <div className='d-flex mb-2'>
            担当者: @{data?.assagin_person}
          </div>
          <div className='d-flex mb-2'>
            状態: {data?.is_complete ? "完了" : "未完了"}
          </div>
          <div className='d-flex mb-2'>
            説明: {data?.description}
          </div>
          <div className='d-flex mb-2'>
            作成日: {new Date(data?.created_at! * 1000).toLocaleDateString() + "," + new Date(data?.created_at! * 1000).toLocaleTimeString()}
          </div>
          <div className='d-flex mb-5'>
            更新日:  {new Date(data?.updated_at! * 1000).toLocaleDateString() + "," + new Date(data?.updated_at! * 1000).toLocaleTimeString()}
          </div>
          <Link to={"/"} className="btn btn-primary btn-lg">戻る</Link>
        </div>
      </div>
    </>
  )
}

export default Detail

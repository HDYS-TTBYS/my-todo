import React, { useEffect } from 'react'
import { Link, useParams, useNavigate } from 'react-router-dom';
import Loading from '../components/Loading';
import NavBar from '../components/NavBar'
import Since from '../components/Since';
import { useMutateTodo } from '../hooks/useMutateTodo';

const Delete = () => {
  const { id } = useParams<"id">();
  const { useMutateDeleteTodo } = useMutateTodo();
  const router = useNavigate();

  useEffect(() => {
    if (useMutateDeleteTodo.isSuccess) { router("/") }
  }, [useMutateDeleteTodo.isSuccess, router])

  if (useMutateDeleteTodo.isLoading) return <Loading />

  return (
    <>
      <NavBar />
      <div className='container'>
        <Since title='削除ページ' />
        <div className='d-flex mt-5 justify-content-evenly'>
          <button type="button" className="btn-lg btn-danger" onClick={() => useMutateDeleteTodo.mutate(Number(id))} >削除</button>
          <Link to={"/"} type="button" className="btn-lg btn-secondary">キャンセル</Link>
        </div>
      </div>
    </>
  )
}

export default Delete

import React from 'react'
import { Link, useParams, useNavigate } from 'react-router-dom';
import NavBar from '../components/NavBar'
import { useMutateTodo } from '../hooks/useMutateTodo';

const Delete = () => {
  const { id } = useParams<"id">();
  const { useMutateDeleteTodo } = useMutateTodo();
  const router = useNavigate();

  if (useMutateDeleteTodo.isSuccess) { router("/") }

  if (useMutateDeleteTodo.isLoading) {
    return <p>Deleting...</p>
  }

  return (
    <>
      <NavBar />
      <div className='container'>
        <div className='d-flex mt-5 justify-content-evenly'>
          <button type="button" className="btn-lg btn-danger" onClick={() => useMutateDeleteTodo.mutate(Number(id))} >削除</button>
          <Link to={"/"} type="button" className="btn-lg btn-secondary">キャンセル</Link>
        </div>
      </div>
    </>
  )
}

export default Delete

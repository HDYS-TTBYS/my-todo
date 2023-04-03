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

  return (
    <>
      <NavBar />
      {useMutateDeleteTodo.isLoading && <Loading />}
      <div className='container'>
        <Since title='Delete Page' />
        <div className='d-flex mt-5 justify-content-evenly'>
          <button type="button" className="btn-lg btn-danger" onClick={() => useMutateDeleteTodo.mutate(Number(id))} >Delete</button>
          <Link to={"/"} type="button" className="btn-lg btn-secondary">Cancel</Link>
        </div>
      </div>
    </>
  )
}

export default Delete

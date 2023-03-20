import React, { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import NavBar from '../components/NavBar'
import { CreateRequestBody, useMutateTodo } from '../hooks/useMutateTodo'

const Create = () => {
  const { useMutateCreateTodo } = useMutateTodo();
  const router = useNavigate();
  const [createTodo, setCreateTodo] = useState<CreateRequestBody>({ title: "", assigin_person: "", description: "" });
  const isOk = createTodo.title !== "" && createTodo.assigin_person !== ""

  if (useMutateCreateTodo.isSuccess) { router("/") }

  if (useMutateCreateTodo.isLoading) {
    return <p>Creating...</p>
  }

  return (
    <>
      <NavBar />
      <div className='container'>
        <div className="mb-3 mt-5">
          <label htmlFor="title" className="form-label">タイトル</label>
          <input type="text" className="form-control" id="title" placeholder="買い物に行く" onChange={(e) => setCreateTodo({ ...createTodo, title: e.target.value })} />
        </div>
        <div className="mb-3">
          <label htmlFor="assigin_person" className="form-label">担当者</label>
          <input type="text" className="form-control" id="assigin_person" placeholder="hdys" onChange={(e) => setCreateTodo({ ...createTodo, assigin_person: e.target.value })} />
        </div>
        <div className="mb-3">
          <label htmlFor="description" className="form-label">説明</label>
          <textarea className="form-control" id="description" rows={3} onChange={(e) => setCreateTodo({ ...createTodo, description: e.target.value })}></textarea>
        </div>
        <div className='d-flex justify-content-evenly'>
          <button type="button" className="btn btn-primary btn-lg" disabled={!isOk} onClick={() => useMutateCreateTodo.mutate(createTodo)}>作成</button>
          <Link to={"/"} type="button" className="btn btn-secondary btn-lg">キャンセル</Link>
        </div>
      </div>
    </>
  )
}

export default Create

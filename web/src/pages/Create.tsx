import React, { useEffect, useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import Loading from '../components/Loading'
import NavBar from '../components/NavBar'
import Since from '../components/Since'
import { CreateRequestBody, useMutateTodo } from '../hooks/useMutateTodo'

const Create = () => {
  const { useMutateCreateTodo } = useMutateTodo();
  const router = useNavigate();
  const [createTodo, setCreateTodo] = useState<CreateRequestBody>({ title: "", assigin_person: "", description: "" });
  const isOk = createTodo.title !== "" && createTodo.assigin_person !== ""

  useEffect(() => {
    if (useMutateCreateTodo.isSuccess) { router("/") }
  }, [useMutateCreateTodo.isSuccess, router])

  return (
    <>
      <NavBar />
      {(useMutateCreateTodo.isLoading) && <Loading />}
      <div className='container'>
        <Since title="Create Page" />
        <div className="mb-3 mt-5">
          <label htmlFor="title" className="form-label">Title</label>
          <input type="text" className="form-control" id="title" placeholder="Title" value={createTodo.title} onChange={(e) => setCreateTodo({ ...createTodo, title: e.target.value })} />
        </div>
        <div className="mb-3">
          <label htmlFor="assigin_person" className="form-label">Assigin Person</label>
          <input type="text" className="form-control" id="assigin_person" placeholder="hdys" value={createTodo.assigin_person} onChange={(e) => setCreateTodo({ ...createTodo, assigin_person: e.target.value })} />
        </div>
        <div className="mb-3">
          <label htmlFor="description" className="form-label">Description</label>
          <textarea className="form-control" id="description" rows={3} value={createTodo.description} onChange={(e) => setCreateTodo({ ...createTodo, description: e.target.value })}></textarea>
        </div>
        <div className='d-flex justify-content-evenly'>
          <button type="button" className="btn btn-primary btn-lg" disabled={!isOk} onClick={() => useMutateCreateTodo.mutate(createTodo)}>Create</button>
          <Link to={"/"} type="button" className="btn btn-secondary btn-lg">Cancel</Link>
        </div>
      </div>
    </>
  )
}

export default Create

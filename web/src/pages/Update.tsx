import React, { useEffect, useState } from 'react'
import { Link, useNavigate, useParams } from 'react-router-dom'
import Loading from '../components/Loading'
import NavBar from '../components/NavBar'
import Since from '../components/Since'
import { RequestBody, useMutateTodo } from '../hooks/useMutateTodo'
import { useQueryTodo } from '../hooks/useQueryTodo'


const Update = () => {
  const { id } = useParams<"id">();
  const { data, isLoading, error } = useQueryTodo(Number(id))
  const { useMutateUpdateTodo } = useMutateTodo();
  const router = useNavigate();
  const [updateTodo, setUpdateTodo] = useState<RequestBody>({ title: "", assigin_person: "", is_complete: false, description: "" });

  useEffect(() => {
    setUpdateTodo({ title: data?.title!, assigin_person: data?.assagin_person!, is_complete: data?.is_complete!, description: data?.description! })
  }, [data])

  const isOk = updateTodo.title !== "" && updateTodo.assigin_person !== ""

  if (isLoading) return <Loading />
  if (useMutateUpdateTodo.isLoading) return <Loading />

  if (useMutateUpdateTodo.isSuccess) { router("/") }
  if (error) { router("/") }
  return (
    <>
      <NavBar />
      <div className='container'>
        <Since title='アップデートページ' />
        <div className="mb-3 mt-5">
          <label htmlFor="title" className="form-label">タイトル</label>
          <input type="text" className="form-control" id="title" placeholder="買い物に行く" value={updateTodo.title} onChange={(e) => setUpdateTodo({ ...updateTodo, title: e.target.value })} />
        </div>
        <div className="mb-3">
          <label htmlFor="assigin_person" className="form-label">担当者</label>
          <input type="text" className="form-control" id="assigin_person" placeholder="hdys" value={updateTodo.assigin_person} onChange={(e) => setUpdateTodo({ ...updateTodo, assigin_person: e.target.value })} />
        </div>
        <div className="form-check form-switch">
          <label className="form-check-label" htmlFor="is_complete">完了</label>
          <input className="form-check-input" type="checkbox" id="is_complete" checked={updateTodo.is_complete} onChange={() => setUpdateTodo({ ...updateTodo, is_complete: !updateTodo.is_complete })} />
        </div>
        <div className="mb-3">
          <label htmlFor="description" className="form-label">説明</label>
          <textarea className="form-control" id="description" value={updateTodo.description} rows={3} onChange={(e) => setUpdateTodo({ ...updateTodo, description: e.target.value })}></textarea>
        </div>
        <div className='d-flex justify-content-evenly'>
          <button type="button" className="btn btn-primary btn-lg" disabled={!isOk} onClick={() => useMutateUpdateTodo.mutate({ id: data?.id!, requestbody: updateTodo })}>更新</button>
          <Link to={"/"} type="button" className="btn btn-secondary btn-lg">キャンセル</Link>
        </div>
      </div>
    </>
  )
}

export default Update

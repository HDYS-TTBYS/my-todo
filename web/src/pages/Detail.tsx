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
        <Since title='Detail Page' />
        <div className='flex mt-5'>
          <div className='d-flex mb-2'>
            Title: {data?.title}
          </div>
          <div className='d-flex mb-2'>
            Assign Person: @{data?.assagin_person}
          </div>
          <div className='d-flex mb-2'>
            Is Complete: {data?.is_complete ? "Completed" : "Not Completed"}
          </div>
          <div className='d-flex mb-2'>
            Description: {data?.description}
          </div>
          <div className='d-flex mb-2'>
            Created At: {new Date(data?.created_at! * 1000).toLocaleDateString() + "," + new Date(data?.created_at! * 1000).toLocaleTimeString()}
          </div>
          <div className='d-flex mb-5'>
            Updated At:  {new Date(data?.updated_at! * 1000).toLocaleDateString() + "," + new Date(data?.updated_at! * 1000).toLocaleTimeString()}
          </div>
          <Link to={"/"} className="btn btn-primary btn-lg">Back</Link>
        </div>
      </div>
    </>
  )
}

export default Detail

import NavBar from '../components/NavBar';
import { useQueryTodos } from '../hooks/useQueryTotos';
import Pagination from "../components/Pagination ";
import Since from '../components/Since';
import TodoItem from '../components/TodoItem';
import { ToDo } from '../generated';
import { Link, useSearchParams } from 'react-router-dom';
import Loading from '../components/Loading';


const Index = () => {
  const [searchParams, setSearchParams] = useSearchParams();
  const offset = Number(searchParams.get("offset"));
  const { data, isLoading, error } = useQueryTodos(offset);

  if (isLoading) return <Loading />

  if (error) return <p>error</p>
  return (
    <>
      <NavBar />
      <div className="container">
        <Since title='Top Page' />

        <div className="my-3 p-3 bg-body rounded shadow-sm">
          <h6 className="border-bottom pb-2 mb-0">Todos</h6>

          <div className="table-responsive">
            <table className='table table-sm'>
              <thead>
                <tr>
                  <th scope="col" className='small'>Detail</th>
                  <th scope="col" className='small'>Completed</th>
                  <th scope="col" className='small'>Update</th>
                  <th scope="col" className='small'>Time</th>
                  <th scope="col" className='small'>Delete</th>
                </tr>
              </thead>
              {data?.todos?.map((todo: ToDo) => (
                <tbody key={todo.id}>
                  <TodoItem todo={todo} />
                </tbody>
              ))}
            </table>
          </div>


        </div>

        <div className='d-flex align-items-center justify-content-between'>
          <Link to={"/create"} type="button" className="btn btn-primary">Create</Link>
          <Pagination totalCount={data?.total!} />
        </div>
      </div>

    </>

  )
}

export default Index

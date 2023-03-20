import NavBar from '../components/NavBar';
import { useQueryTodos } from '../hooks/useQueryTotos';
import Pagination from "../components/Pagination ";
import Since from '../components/Since';
import TodoItem from '../components/TodoItem';
import { ToDo } from '../generated';
import { Link } from 'react-router-dom';
import Loading from '../components/Loading';


const Index = () => {
  const { data, isLoading, error } = useQueryTodos(0)

  if (isLoading) return <Loading />

  if (error) return <p>error</p>
  return (
    <>
      <NavBar />
      <div className="container">
        <Since title='トップページ' />

        <div className="my-3 p-3 bg-body rounded shadow-sm">
          <h6 className="border-bottom pb-2 mb-0">Todos</h6>

          {data?.todos?.map((todo: ToDo) => (
            <div key={todo.id}>
              <TodoItem todo={todo} />
            </div>
          ))}

        </div>

        <div className='d-flex align-items-center justify-content-between'>
          <Link to={"/create"} type="button" className="btn btn-primary btn-lg">新規作成</Link>
          <Pagination />
        </div>
      </div>

    </>

  )
}

export default Index

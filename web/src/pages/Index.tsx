import NavBar from '../components/NavBar';
import { useQueryTodos } from '../hooks/useQueryTotos';
import Pagination from "../components/Pagination ";
import Since from '../components/Since';
import TodoItem from '../components/TodoItem';


const Index = () => {
  const { data, isLoading, error } = useQueryTodos(0)

  if (isLoading) return <div> 'Loading...'</div>

  if (error) return <div>'An error has occurred: ' + error</div>
  return (
    <>

      <NavBar />
      <div className="container">
        <Since />

        <div className="my-3 p-3 bg-body rounded shadow-sm">
          <h6 className="border-bottom pb-2 mb-0">Todos</h6>

          {data?.todos?.map((todo) => (
            <div key={todo.id}>
              <div>{todo.title}
              </div>
              <TodoItem todo={todo} />
            </div>
          ))}

        </div>

        <Pagination />
      </div>

    </>

  )
}

export default Index

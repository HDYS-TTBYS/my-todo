import NavBar from '../components/NavBar';
import { useQueryTodos } from '../hooks/useQueryTotos';

const Index = () => {
  const { data } = useQueryTodos(0)

  return (
    <div>
      <NavBar />
      {JSON.stringify(data)}
    </div>
  )
}

export default Index

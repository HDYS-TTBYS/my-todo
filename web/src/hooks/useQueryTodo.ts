import { useQuery } from "react-query";
import { ToDo } from "../types/typescript-fetch/models";

export const useQueryTodo = (id: number) => {
  const getTodo = async () => {
    const res = await fetch(`${process.env.REACT_APP_REST_URL}/api/todo/${id}`)
    if (res.ok) {
      return res.json() as Promise<ToDo>
    }
    return res.json() as Promise<Error>
  }

  return useQuery({
    queryKey: "todo",
    queryFn: getTodo,
  })
}

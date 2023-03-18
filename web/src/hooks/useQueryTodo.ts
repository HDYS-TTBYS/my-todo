import { useQuery } from "react-query";
import axios from "axios";
import { ToDo } from "../types/typescript-fetch/models";

export const useQueryTodo = (id: number) => {
  const getTodo = async () => {
    const data = await axios.get<ToDo>(
      `${process.env.REACT_APP_REST_URL}/api/todo/${id}`
    )
    return data
  }

  return useQuery({
    queryKey: "todo",
    queryFn: getTodo,
  })
}

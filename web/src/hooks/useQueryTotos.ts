import { useQuery } from "react-query";
import axios from "axios";
import { GetTodos200Response } from "../types/typescript-fetch";

export const useQueryTodos = (offset: number) => {
  const getTodos = async () => {
    const data = await axios.get<GetTodos200Response>(
      `${process.env.REACT_APP_REST_URL}/api/todos?offset=${offset}`
    )
    console.log(data)
    return data.data
  }

  return useQuery({
    queryKey: "todos",
    queryFn: getTodos,
  })
}

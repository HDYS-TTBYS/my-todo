import { useQueryClient, useMutation } from "react-query";
import appClient from "../app-client";

export type CreateRequestBody = {
  title: string;
  description: string;
  assigin_person: string;
}

export type UpdateRequestBody = {
  id: number;
  requestbody: RequestBody
}
export type RequestBody = {
  title: string;
  description: string;
  assigin_person: string;
  is_complete: boolean;
}

export const useMutateTodo = () => {
  const queryClient = useQueryClient()

  const useMutateCreateTodo = useMutation(
    (requestBody: CreateRequestBody) =>
      appClient.todo.postTodo(requestBody),
    {
      onSuccess: (res) => {
        queryClient.invalidateQueries('todos')
      },
    }
  )

  const useMutateUpdateTodo = useMutation(
    (requestBody: UpdateRequestBody) =>
      appClient.todo.updateTodoId(requestBody.id, requestBody.requestbody),
    {
      onSuccess: (res) => {
        queryClient.invalidateQueries('todos')
      },
    }
  )

  const useMutateDeleteTodo = useMutation(
    (id: number) =>
      appClient.todo.deleteTodoId(id),
    {
      onSuccess: (res) => {
        queryClient.invalidateQueries('todos')
      },
    }
  )

  return { useMutateCreateTodo, useMutateUpdateTodo, useMutateDeleteTodo }
}

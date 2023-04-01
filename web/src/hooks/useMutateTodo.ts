import { useAtom } from "jotai";
import { useQueryClient, useMutation } from "react-query";
import appClient from "../app-client";
import { ApiError } from "../generated";
import { alertAtom } from "../store";

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

const ALERT_TIMER = 3

export const useMutateTodo = () => {
  const queryClient = useQueryClient();
  const [, setAlert] = useAtom(alertAtom);

  const useMutateCreateTodo = useMutation(
    (requestBody: CreateRequestBody) =>
      appClient.todo.postTodo(requestBody),
    {
      onSuccess: (res) => {
        setAlert({ alert: "success", message: "todo is created" });
        setTimeout(() => {
          setAlert({ alert: "success", message: "" });
        }, ALERT_TIMER * 1000);
        queryClient.invalidateQueries('todos')
      },
      onError: (error: ApiError) => {
        setAlert({ alert: "danger", message: error.body.message });
        setTimeout(() => {
          setAlert({ alert: "danger", message: "" });
        }, ALERT_TIMER * 1000);
      },
    }
  )

  const useMutateUpdateTodo = useMutation(
    (requestBody: UpdateRequestBody) =>
      appClient.todo.updateTodoId(requestBody.id, requestBody.requestbody),
    {
      onSuccess: (res) => {
        setAlert({ alert: "success", message: "todo is updated" });
        setTimeout(() => {
          setAlert({ alert: "success", message: "" });
        }, ALERT_TIMER * 1000);
        queryClient.invalidateQueries('todos')
      },
      onError: (error: ApiError) => {
        setAlert({ alert: "danger", message: error.body.message });
        setTimeout(() => {
          setAlert({ alert: "danger", message: "" });
        }, ALERT_TIMER * 1000);
      },
    }
  )

  const useMutateDeleteTodo = useMutation(
    (id: number) =>
      appClient.todo.deleteTodoId(id),
    {
      onSuccess: (res) => {
        setAlert({ alert: "success", message: "todo is deleted" });
        setTimeout(() => {
          setAlert({ alert: "success", message: "" });
        }, ALERT_TIMER * 1000);
        queryClient.invalidateQueries('todos')
      },
      onError: (error: ApiError) => {
        setAlert({ alert: "danger", message: error.body.message });
        setTimeout(() => {
          setAlert({ alert: "danger", message: "" });
        }, ALERT_TIMER * 1000);
      },
    }
  )

  return { useMutateCreateTodo, useMutateUpdateTodo, useMutateDeleteTodo }
}

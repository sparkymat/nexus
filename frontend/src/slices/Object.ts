import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { Object } from '../models/Object';

export interface CreateRequest {
  name: string;
  isTemplate: boolean;
  templateId?: string;
}

export const api = createApi({
  reducerPath: 'objects',
  tagTypes: ['Object'],
  baseQuery: fetchBaseQuery({ baseUrl: '/api/' }),
  endpoints: builder => ({
    create: builder.mutation<Object, CreateRequest>({
      query: body => ({
        url: 'objects',
        method: 'POST',
        body,
        headers: {
          'X-CSRF-Token': (
            document.querySelector('meta[name="csrf-token"]') as any
          ).content,
        },
      }),
      invalidatesTags: [{ type: 'Object', id: 'LIST' }],
    }),
    show: builder.query<Object, string>({
      query: id => `objects/${id}`,
      providesTags: (_result, _error, arg) => [{ type: 'Object', id: arg }],
    }),
  }),
});

export const { useCreateMutation, useShowQuery } = api;

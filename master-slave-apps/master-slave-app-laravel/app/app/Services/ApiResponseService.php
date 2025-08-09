<?php

namespace App\Services;

use Illuminate\Http\JsonResponse;

class ApiResponseService
{
    /**
     * Mengembalikan respons JSON dengan struktur standar
     *
     * @param int $code
     * @param string $message
     * @param mixed $data
     * @return JsonResponse
     */
    public function send(int $code, string $message, $data = null, $meta = null): JsonResponse
    {
        $response = [
            'code' => $code,
            'message' => $message,
            'data' => $data,
        ];


        if (!is_null($meta)) {
            $response['meta'] = $meta;
        }

        return response()->json($response, $code);
    }

    /**
     * Respons sukses
     *
     * @param string $message
     * @param mixed $data
     * @param int $code
     * @return JsonResponse
     */
    public function success(int $code = 200, string $message = 'success', $data = null): JsonResponse
    {
        return $this->send($code, $message, $data);
    }

    /**
     * Respons error
     *
     * @param string $message
     * @param int $code
     * @return JsonResponse
     */
    public function error(int $code = 400, string $message = 'Error'): JsonResponse
    {
        return $this->send($code, $message);
    }

    public function validationError(int $code = 400,string $message = 'Error',  $errValidation = null): JsonResponse
    {
        $response = [
            'code' => $code,
            'message' => $message,
            'errors' => $errValidation,
        ];

        return response()->json($response, $code);
    }
}
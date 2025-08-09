<?php

use Illuminate\Foundation\Application;
use Illuminate\Foundation\Configuration\Exceptions;
use Illuminate\Foundation\Configuration\Middleware;
use Illuminate\Http\JsonResponse;
use Symfony\Component\HttpFoundation\Response;


return Application::configure(basePath: dirname(__DIR__))
    ->withRouting(
        web: __DIR__.'/../routes/web.php',
        commands: __DIR__.'/../routes/console.php',
        health: '/up',
    )
    ->withMiddleware(function (Middleware $middleware) {
        $middleware->validateCsrfTokens(except: [
            'items', // <-- exclude this route,
            'items/*' // <-- exclude this route,
        ]);
    })
    ->withExceptions(function (Exceptions $exceptions) {
        $exceptions->respond(function (Response $response) {
            if ($response->getStatusCode() === 405) {
                return response()->json([
                    'code' => 404,
                    'message' => 'Rute tidak ditemukan',
                    'data' => null,
                    
                ], 404);
            }
    
            return $response;
        });
    })->create();
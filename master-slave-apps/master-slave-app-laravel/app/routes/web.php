<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\ItemController;

Route::get('/', function () {
    return view('welcome');
});


Route::controller(ItemController::class)->group(function () {
    Route::get('/items', 'index');
    Route::post('/items', 'store');
    Route::get('/items/{id}', 'show');
    Route::put('/items/{id}', 'update');
    Route::delete('/items/{id}', 'destroy');
});
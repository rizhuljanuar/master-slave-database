<?php

namespace App\Http\Controllers;

use App\Models\Item;
use Illuminate\Http\Request;
use App\Services\ApiResponseService;
use Illuminate\Support\Facades\Validator;

class ItemController extends Controller
{

    protected $customResponse;

    public function __construct(ApiResponseService $customResponse)
    {
        $this->customResponse = $customResponse;
    }


    // Menampilkan semua data (baca dari Slave)
    public function index()
    {
        $items = Item::paginate(10);
       
         if ($items->isEmpty()){
            return $this->customResponse->send(404, 'Item tidak tersedia', null);
        }
        
        return $this->customResponse->send(200, 'success', $items->items(), null);
    }

    // Menampilkan data berdasarkan ID (baca dari Slave)
    public function show($id)
    {
        $item = Item::find($id);

        if (empty($item)){
            return $this->customResponse->send(404, 'Item tidak ditemukan', null);
        }
        
        return $this->customResponse->send(200, 'success', $item);
    }

    // Menyimpan data baru (tulis ke Master)
    public function store(Request $request)
    {
    
        $validator = Validator::make($request->all(), [
            'title' => 'required|string|max:150',
            'description' => 'nullable|string',
            'qty' => 'required|integer|min:0',
            'price' => 'required|numeric|min:1',
        ]);

        if ($validator->fails()) {
            return $this->customResponse->validationError(422, 'validation error', $validator->errors());
        }


        $item = Item::create([
            'title' => $request->title,
            'description' => $request->description,
            'qty' => $request->qty,
            'price' => $request->price,
        ]);

        $itemS = Item::find($item->id);

        $result  =  [
            'id' => $itemS->id,
            'title' => $request->title,
            'description' => $request->description,
            'qty' => $request->qty,
            'price' => $request->price,
            'created_at' => $itemS->created_at,
        ];
    
        return $this->customResponse->send(200, 'success', $result);
        
    }

    // Memperbarui data (tulis ke Master)
    public function update(Request $request, $id)
    {
        $item = Item::find($id);

         if (empty($item)){
            return $this->customResponse->send(404, 'Item tidak ditemukan', null);
        }

        $validator = Validator::make($request->all(), [
            'title' => 'sometimes|string|max:255',
            'description' => 'sometimes|nullable|string',
            'qty' => 'sometimes|integer|min:0',
            'price' => 'sometimes|numeric|min:0',
        ]);

        if ($validator->fails()) {
            return $this->customResponse->validationError(422, 'validation error', $validator->errors());
        }



        $item->update([
            'title' => $request->title,
            'description' => $request->description,
            'qty' => $request->qty,
            'price' => $request->price,
            'updated_at' => date('Y-m-d H:i:s'),
        ]);
        return $this->customResponse->send(200, 'Item berhasil di update', $item);
    }

    // Menghapus data (tulis ke Master)
    public function destroy($id)
    {
         $item = Item::find($id);

         if (empty($item)){
            return $this->customResponse->send(404, 'Item tidak ditemukan', null);
        }

        $item->delete();
        return $this->customResponse->send(200, 'Item berhasil dihapus', null);
    }
}
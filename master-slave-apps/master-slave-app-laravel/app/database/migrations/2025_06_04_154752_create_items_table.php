<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        if (!Schema::hasTable('items')) {
            Schema::create('items', function (Blueprint $table) {
                $table->id();
                $table->string('title',150)->default('-')->index();
                $table->longText('description');
                $table->unsignedTinyInteger('qty');
                $table->decimal('price', 12, 2); // 8 digit total, 2 digit desimal (misalnya: 123456.78)
                $table->timestamp('created_at')->useCurrent()->index('created_at_index')->comment('Waktu pembuatan');
                $table->timestamp('updated_at')->nullable()->comment('Waktu perubahan, boleh null');
            });
        }
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('items');
    }
};

<?php
namespace Tests\Unit;

use Tests\TestCase;
use App\Models\Product;
use App\Services\ProductService;
use Mockery;
use Illuminate\Foundation\Testing\RefreshDatabase;

class ProductControllerTest extends TestCase
{
    use RefreshDatabase;

    protected $productService;

    public function setUp(): void
    {
        parent::setUp();
        $this->productService = Mockery::mock(ProductService::class);
        $this->app->instance(ProductService::class, $this->productService);
    }

    public function tearDown(): void
    {
        Mockery::close();
        parent::tearDown();
    }

    public function testIndex()
    {
        // Mock the 'all' method from the ProductService to return a collection of Product objects
        $this->productService
            ->shouldReceive('all')
            ->once();

        // Perform a GET request to the '/api/products' endpoint
        $response = $this->get('/api/products');

        // Check if the response status is 200 OK
        $response->assertStatus(200)
            ->assertJsonStructure([
                '*' => [  // Use '*' to indicate that the structure should apply to all elements in the array
                    'id', 'name', 'description', 'price', 'stock', 'created_at', 'updated_at'
                ]
            ]);
    }

    public function testShow()
    {
        $product = Product::factory()->create();

        $this->productService
            ->shouldReceive('find')
            ->with($product->id)
            ->once()
            ->andReturn($product);

        $response = $this->get('/api/products/' . $product->id);

        $response->assertStatus(200)
            ->assertJson([
                'id' => $product->id,
                'name' => $product->name,
                'description' => $product->description,
                'price' => $product->price,
                'stock' => $product->stock
            ]);
    }

    public function testStore()
    {
        $data = [
            'name' => 'Test Product',
            'description' => 'Test Description',
            'price' => 100,
            'stock' => 10
        ];

        $this->productService
            ->shouldReceive('create')
            ->once()
            ->andReturn((new Product())->forceFill($data));

        $response = $this->post('/api/products', $data);

        $response->assertStatus(201)
            ->assertJson($data);
    }

    public function testUpdate()
    {
        $product = Product::factory()->create();

        $data = [
            'name' => 'Updated Product',
            'description' => 'Updated Description',
            'price' => 150,
            'stock' => 20
        ];

        $this->productService
            ->shouldReceive('update')
            ->with($product->id, $data)
            ->once()
            ->andReturn((new Product())->forceFill(array_merge(['id' => $product->id], $data)));

        $response = $this->put('/api/products/' . $product->id, $data);

        $response->assertStatus(200)
            ->assertJson($data);
    }

    public function testDestroy()
    {
        $product = Product::factory()->create();

        $this->productService
            ->shouldReceive('delete')
            ->with($product->id)
            ->once()
            ->andReturn(true);

        $response = $this->delete('/api/products/' . $product->id);

        $response->assertStatus(200);
    }
}


import Drill from "$lib/assets/categories/drill.png";
import Pump from "$lib/assets/categories/pump.png";
import Generator from "$lib/assets/categories/generator.png";
import Forklift from "$lib/assets/categories/forklift.png";
import Van from "$lib/assets/categories/van.png";

interface ProductCategory {
    id: string;
    translationKey: string;
    image: string;
}

export interface ProductConfig {
    categories?: ProductCategory[],
    categoryFallback: ProductCategory
}

export const ProductConfigKey = 'product';

export const defaultProductConfig: { [ProductConfigKey]: ProductConfig } = {
    [ProductConfigKey]: {
        categoryFallback: {
            id: "other",
            translationKey: "product.category.other",
            image: "https://via.placeholder.com/150"
        },
        categories: [
            {
                id: "drill",
                translationKey: "product.category.drill",
                image: Drill
            },
            {
                id: "pump",
                translationKey: "product.category.pump",
                image: Pump
            },
            {
                id: "generator",
                translationKey: "product.category.generator",
                image: Generator
            },
            {
                id: "forklift",
                translationKey: "product.category.forklift",
                image: Forklift
            },
            {
                id: "van",
                translationKey: "product.category.van",
                image: Van
            }
        ]
    }
}

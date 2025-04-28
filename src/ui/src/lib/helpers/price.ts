export function formatCurrency(price: number, currency: string) {
    return (price / 100).toFixed(2) + " " + currency;
}

export function formatStringToCurrency(price: string): number {
    return Number((parseFloat(price) * 100).toFixed(0));
}
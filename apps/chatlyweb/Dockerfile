FROM oven/bun:alpine AS builder

WORKDIR /app

COPY . .

RUN bun install --frozen-lockfile

RUN bun run build

FROM oven/bun:alpine AS runner

WORKDIR /app

ENV NODE_ENV=production
ENV PORT=3000

COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next/static ./.next/static

EXPOSE 3000

CMD ["bun", "server.js"]
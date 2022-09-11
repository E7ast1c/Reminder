create table if not exists public.reminders
(
    id          serial primary key,
    description text      not null,
    chat_id     bigint   not null,
    created_at     timestamp not null,
    last_due    timestamp not null,
    scheduler   text
);
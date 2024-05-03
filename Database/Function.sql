
CREATE OR REPLACE FUNCTION public.f_check_session_refresh_token(ref_token character varying)
 RETURNS TABLE(is_not_expired boolean, customer_id uuid, remember_me boolean, id_session uuid, group_api jsonb)
 LANGUAGE plpgsql
AS $function$

BEGIN 

	return query 
	select
	case 
		when expired_at >= now() then true
		else false
	end as is_not_expired,
	s.customer_id,
	s.remember_me,
	s.id as id_session,
	s.group_api 
	from
		public.sessions s
	where
		(encode(public.hmac(id::text, key, 'sha256'), 'hex')) = REF_TOKEN;

END;

$function$
;

CREATE OR REPLACE FUNCTION public.f_check_user(user_name1 character varying)
 RETURNS TABLE(customer_id uuid, is_true boolean)
 LANGUAGE plpgsql
AS $function$
BEGIN 
	return query 
	select
		c.customer_id,
		case
			when COUNT(c.customer_id) > 0 then true
			else false
		end as is_true
	from
		public.customer c
	where (c.username = USER_NAME1)
	group by c.customer_id ;
END;
$function$
;



CREATE OR REPLACE FUNCTION public.f_create_session_refresh_token_user(user_id1 character varying, remember_me1 boolean, interval_month_remember1 character varying, interval_day_remember1 character varying, groupapi1 character varying)
 RETURNS TABLE(refresh_token text)
 LANGUAGE plpgsql
AS $function$
BEGIN 
	return query 
		with in_sessions as (
		insert
			into
				public.sessions (
					customer_id,
					remember_me,
					expired_at
				)
				select
					user_id1::uuid,
					REMEMBER_ME1,
					case
						when REMEMBER_ME1 = true then now()+ INTERVAL_MONTH_REMEMBER1::interval
						else now()+ INTERVAL_DAY_REMEMBER1::interval
					end
					returning *)
		select
			(encode(public.hmac(id::text, 'key', 'sha256'), 'hex')) as refresh_token
		from
			in_sessions;
END;
$function$
;


CREATE OR REPLACE FUNCTION public.f_get_all_product(category_id1 character varying)
 RETURNS TABLE(product_id uuid, name character varying, stock integer, price numeric, category text)
 LANGUAGE plpgsql
AS $function$
	begin
		return query
		select 
			p.product_id,
			p."name",
			p.stock,
			p.price,
			STRING_AGG(c.category_name::text, ', ') AS Category
		from public.products p 
		join public.product_category pc on pc.product_id = p.product_id
		join public.category c on c.category_id = pc.category_id
		where p.is_deleted is false
		and (case when category_id1 is not null then pc.category_id=category_id1::uuid else true end)
		GROUP by p.product_id
		order by p."name" asc;

	END;
$function$
;

CREATE OR REPLACE FUNCTION public.f_get_cart(customer_id1 character varying)
 RETURNS TABLE(cart_id uuid, name character varying, quantity integer)
 LANGUAGE plpgsql
AS $function$
	begin
		return query
		select
		c.cart_id,
		p."name",
		pc.quantity
		from public.carts c 
		join public.product_carts pc on pc.cart_id = c.cart_id 
		join public.products p on p.product_id = pc.product_id 
		where c.is_deleted is false 
		and c.customer_id = customer_id1::uuid;
	END;
$function$
;


CREATE OR REPLACE FUNCTION public.f_get_log_data_user(user_id1 character varying)
 RETURNS TABLE(customer_id character varying, customer_email character varying, username character varying, created_at character varying)
 LANGUAGE plpgsql
AS $function$
BEGIN 
	return query(
		select
			c.customer_id::varchar,
			c.customer_email,
			c.username,
			c.created_at::varchar
		from
			public.customer c
		where
			c.customer_id = USER_ID1::uuid
			limit 1);
END;
$function$
;

CREATE OR REPLACE FUNCTION public.f_get_password_user(user_id1 character varying)
 RETURNS TABLE(credential character varying)
 LANGUAGE plpgsql
AS $function$
BEGIN 
	return query 
	select c.credential from public.customer c WHERE c.customer_id=user_id1
::uuid;
END;
$function$
;


CREATE OR REPLACE FUNCTION public.f_update_session(ref_token character varying, time_str character varying)
 RETURNS character varying
 LANGUAGE plpgsql
AS $function$
BEGIN 
	update public.sessions set expired_at=now()+ TIME_STR::interval
	where (encode(public.hmac(id::text, key, 'sha256'), 'hex')) = REF_TOKEN;
	return 'ok';
END;
$function$
;


CREATE OR REPLACE PROCEDURE public.p_add_cart(IN customer_id1 character varying, IN product_id1 character varying, IN quantity1 integer)
 LANGUAGE plpgsql
AS $procedure$
DECLARE 
    cart_ids uuid;
    count_filter_cart int;
    count_filter_cart_item int;
BEGIN
    SELECT c.cart_id 
    INTO cart_ids
    FROM public.carts c 
    WHERE c.customer_id = customer_id1::uuid 
    AND c.is_deleted is false;
 	
    IF cart_ids is null then
        INSERT INTO public.carts (customer_id)
        VALUES (customer_id1::uuid)
        RETURNING cart_id INTO cart_ids;
        
        INSERT INTO public.product_carts (product_id, cart_id, quantity)
        VALUES (product_id1::uuid, cart_ids, quantity1);
    ELSE 
        SELECT COUNT(pc.product_cart_id)
        INTO count_filter_cart_item
        FROM public.product_carts pc 
        WHERE pc.product_id = product_id1::uuid
        AND pc.cart_id = cart_ids;
        IF count_filter_cart_item <= 0 THEN 
            UPDATE public.product_carts 
            SET quantity = quantity + quantity1
            WHERE product_id = product_id1::uuid
            AND cart_id = cart_ids;
        ELSE 
        
            INSERT INTO public.product_carts (product_id, cart_id, quantity)
            VALUES (product_id1::uuid, cart_ids, quantity1);
        END IF;
    END IF;
    
    RAISE INFO '%', 'Sucessfuly';
END;
$procedure$
;

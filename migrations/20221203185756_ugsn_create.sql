-- +goose Up
-- +goose StatementBegin
create table if not exists ugsn(
    code varchar(10) primary key not null,
    title varchar(500) not null
);

insert into ugsn(code, title) values ('01.00.00', 'Математика и механика');
insert into ugsn(code, title) values ('02.00.00', 'Компьютерные и информационные науки');
insert into ugsn(code, title) values ('04.00.00', 'Химия');
insert into ugsn(code, title) values ('03.00.00', 'Физика и астрономия');
insert into ugsn(code, title) values ('05.00.00', 'Науки о земле');
insert into ugsn(code, title) values ('06.00.00', 'Биологические науки');
insert into ugsn(code, title) values ('07.00.00', 'Архитектура');
insert into ugsn(code, title) values ('08.00.00', 'Техника и технологии строительства');
insert into ugsn(code, title) values ('09.00.00', 'Информатика и вычислительная техника');
insert into ugsn(code, title) values ('10.00.00', 'Информационная безопасность');
insert into ugsn(code, title) values ('11.00.00', 'Электроника, радиотехника и системы связи');
insert into ugsn(code, title) values ('12.00.00', 'Фотоника, приборостроение, оптические и биотехнические системы и технологии');
insert into ugsn(code, title) values ('13.00.00', 'Электро- и теплоэнергетика');
insert into ugsn(code, title) values ('14.00.00', 'Ядерная энергетика и технологии');
insert into ugsn(code, title) values ('15.00.00', 'Машиностроение');
insert into ugsn(code, title) values ('16.00.00', 'Физико-технические науки и технологии');
insert into ugsn(code, title) values ('17.00.00', 'Оружие и системы вооружения');
insert into ugsn(code, title) values ('18.00.00', 'Химические технологии');
insert into ugsn(code, title) values ('19.00.00', 'Промышленная экология и биотехнологии');
insert into ugsn(code, title) values ('20.00.00', 'Техносферная безопасность и природообустройство');
insert into ugsn(code, title) values ('21.00.00', 'Прикладная геология, горное дело, нефтегазовое дело и геодезия');
insert into ugsn(code, title) values ('22.00.00', 'Технологии материалов');
insert into ugsn(code, title) values ('23.00.00', 'Техника и технологии наземного транспорта');
insert into ugsn(code, title) values ('24.00.00', 'Авиационная и ракетно-космическая техника');
insert into ugsn(code, title) values ('25.00.00', 'Аэронавигация и эксплуатация авиационной и ракетно-космической техники');
insert into ugsn(code, title) values ('26.00.00', 'Техника и технологии кораблестроения и водного транспорта');
insert into ugsn(code, title) values ('27.00.00', 'Управление в технических системах');
insert into ugsn(code, title) values ('28.00.00', 'Нанотехнологии и материалы');
insert into ugsn(code, title) values ('29.00.00', 'Технологии легкой промышленности');
insert into ugsn(code, title) values ('30.00.00', 'Фундаментальная медицина');
insert into ugsn(code, title) values ('31.00.00', 'Клиническая медицина');
insert into ugsn(code, title) values ('32.00.00', 'Науки о здоровье и профилактическая медицина');
insert into ugsn(code, title) values ('33.00.00', 'Фармация');
insert into ugsn(code, title) values ('34.00.00', 'Сестринское дело');
insert into ugsn(code, title) values ('35.00.00', 'Сельское, лесное и рыбное хозяйство');
insert into ugsn(code, title) values ('36.00.00', 'Ветеринария и зоотехния');
insert into ugsn(code, title) values ('37.00.00', 'Психологические науки');
insert into ugsn(code, title) values ('38.00.00', 'Экономика и управление');
insert into ugsn(code, title) values ('39.00.00', 'Социология и социальная работа');
insert into ugsn(code, title) values ('40.00.00', 'Юриспруденция');
insert into ugsn(code, title) values ('41.00.00', 'Политические науки и регионоведение');
insert into ugsn(code, title) values ('42.00.00', 'Средства массовой информации и информационно-библиотечное дело');
insert into ugsn(code, title) values ('43.00.00', 'Сервис и туризм');
insert into ugsn(code, title) values ('44.00.00', 'Образование и педагогические науки');
insert into ugsn(code, title) values ('45.00.00', 'Языкознание и литературоведение');
insert into ugsn(code, title) values ('46.00.00', 'История и археология');
insert into ugsn(code, title) values ('47.00.00', 'Философия, этика и религиоведение');
insert into ugsn(code, title) values ('48.00.00', 'Теология');
insert into ugsn(code, title) values ('49.00.00', 'Физическая культура и спорт');
insert into ugsn(code, title) values ('50.00.00', 'Искусствознание');
insert into ugsn(code, title) values ('51.00.00', 'Культуроведение и социокультурные проекты');
insert into ugsn(code, title) values ('52.00.00', 'Сценические искусства и литературное творчество');
insert into ugsn(code, title) values ('53.00.00', 'Музыкальное искусство');
insert into ugsn(code, title) values ('54.00.00', 'Изобразительное и прикладные виды искусств');
insert into ugsn(code, title) values ('55.00.00', 'Экранные искусства');
insert into ugsn(code, title) values ('58.00.00', 'Востоковедение и африканистика');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists ugsn;
-- +goose StatementEnd

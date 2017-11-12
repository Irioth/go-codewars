package codewars

/**
 * Возможные действия игрока.
 * <p>
 * Игрок не может совершить новое действие, если в течение последних {@code game.actionDetectionInterval - 1} игровых
 * тиков он уже совершил максимально возможное для него количество действий. В начале игры это ограничение для каждого
 * игрока равно {@code game.BaseActionCount}. Ограничение увеличивается за каждый контролируемый игроком центр
 * управления ({@code FacilityType.CONTROL_CENTER}).
 * <p>
 * Большинство действий требует указания дополнительных параметров, являющихся полями объекта {@code move}. В случае,
 * если эти параметры установлены некорректно либо указаны не все обязательные параметры, действие будет проигнорировано
 * игровым симулятором. Любое действие, отличное от {@code NONE}, даже проигнорированное, будет учтено в счётчике
 * действий игрока.
 */

type ActionType byte

const (
	/**
	 * Ничего не делать.
	 */
	Action_None ActionType = iota

	/**
	 * Пометить юнитов, соответствующих некоторым параметрам, как выделенных.
	 * При этом, со всех остальных юнитов выделение снимается.
	 * Юниты других игроков автоматически исключаются из выделения.
	 */
	Action_ClearAndSelect

	/**
	 * Пометить юнитов, соответствующих некоторым параметрам, как выделенных.
	 * При этом, выделенные ранее юниты остаются выделенными.
	 * Юниты других игроков автоматически исключаются из выделения.
	 */
	Action_AddToSelection

	/**
	 * Снять выделение с юнитов, соответствующих некоторым параметрам.
	 */
	Action_Deselect

	/**
	 * Установить для выделенных юнитов принадлежность к группе.
	 */
	Action_Assign

	/**
	 * Убрать у выделенных юнитов принадлежность к группе.
	 */
	Action_Dismiss

	/**
	 * Расформировать группу.
	 */
	Action_Disband

	/**
	 * Приказать выделенным юнитам меремещаться в указанном направлении.
	 */
	Action_Move

	/**
	 * Приказать выделенным юнитам поворачиваться относительно некоторой точки.
	 */
	Action_Rotate

	/**
	 * Масштабировать формацию выделенных юнитов относительно указанной точки.
	 */
	Action_Scale

	/**
	 * Настроить производство нужного типа техники на заводе ({@code FacilityType.VEHICLE_FACTORY}).
	 */
	Action_SetupVehicleProduction

	/**
	 * Запросить тактический ядерный удар.
	 */
	Action_TacticalNuclearStrike
)

/**
 * Тип местности.
 */
type Terrain byte

const (
	/**
	 * Равнина.
	 */
	Terrain_Plain Terrain = iota

	/**
	 * Топь.
	 */
	Terrain_Swamp

	/**
	 * Лес.
	 */
	Terrain_Forest
)

/**
 * Тип погоды.
 */
type Weather byte

const (
	/**
	 * Ясно.
	 */
	Weather_Clear Weather = iota

	/**
	 * Плотные облака.
	 */
	Weather_Cloud

	/**
	 * Сильный дождь.
	 */
	Weather_Rain
)

/**
 * Тип техники.
 */
type VehicleType uint8

const Vehicle_None VehicleType = 255

const (
	/**
	 * Бронированная ремонтно-эвакуационная машина. Наземный юнит.
	 * Постепенно восстанавливает прочность находящейся поблизости неподвижной техники.
	 */
	Vehicle_Arrv VehicleType = iota

	/**
	 * Истребитель. Воздушный юнит. Крайне эффективен против другой воздушной техники. Не может атаковать наземные цели.
	 */
	Vehicle_Fighter

	/**
	 * Ударный вертолёт. Воздушный юнит. Может атаковать как воздушные, так и наземные цели.
	 */
	Vehicle_Helicopter

	/**
	 * Боевая машина пехоты. Наземный юнит. Может атаковать как воздушные, так и наземные цели.
	 */
	Vehicle_Ifv

	/**
	 * Танк. Наземный юнит. Крайне эффективен против другой наземной техники. Также может атаковать воздушные цели.
	 */
	Vehicle_Tank
)

/**
 * Тип сооружения.
 */
type FacilityType byte

const (
	/**
	 * Центр управления. Увеличивает возможное количество действий игрока на
	 * {@code game.AdditionalActionCountPerControlCenter} за {@code game.actionDetectionInterval} игровых тиков.
	 */
	Facility_ControlCenter FacilityType = iota

	/**
	 * Завод. Может производить технику любого типа по выбору игрока.
	 */
	Facility_VehicleFactory
)
